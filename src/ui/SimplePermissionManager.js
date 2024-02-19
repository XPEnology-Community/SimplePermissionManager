// Namespace definition
Ext.ns("SynoCommunity.SimplePermissionManager");

// Application definition
Ext.define("SynoCommunity.SimplePermissionManager.AppInstance", {
    extend: "SYNO.SDS.AppInstance",
    appWindowName: "SynoCommunity.SimplePermissionManager.AppWindow",
    constructor: function () {
        this.callParent(arguments);
    },
});

// Translator
_V = function (category, element) {
    return _TT("SynoCommunity.SimplePermissionManager.AppInstance", category, element)
}

// Window definition
Ext.define("SynoCommunity.SimplePermissionManager.AppWindow", {
    extend: "SYNO.SDS.AppWindow",
    appInstance: null,
    tabs: null,
    onOpen: function (a) {
        SynoCommunity.SimplePermissionManager.AppWindow.superclass.onOpen.call(
            this,
            a
        );
    },
    constructor: function (config) {
        this.appInstance = config.appInstance;

        this.tabs = function () {
            var allTabs = [];

            allTabs.push({
                title: _V("ui", "overview"),
                items: [
                    this.createDisplayStatus(),
                    this.createDisplayConfig(),
                ],
            });

            allTabs.push({
                title: _V("ui", "packages"),
                layout: "fit",
                items: [this.createPackageGrid()],
            });

            allTabs.push({
                title: _V("ui", "users"),
                layout: "fit",
                items: [this.createUserGrid()],
            });

            return allTabs;
        }.call(this);

        config = Ext.apply(
            {
                resizable: true,
                maximizable: true,
                minimizable: true,
                width: 640,
                height: 510,
                padding: "15px",
                items: [
                    {
                        xtype: "syno_tabpanel",
                        activeTab: 0,
                        plain: true,
                        items: this.tabs,
                        deferredRender: true,
                    },
                ],
            },
            config
        );

        this.callParent([config]);
    },
    // Create the display of status
    createDisplayStatus: function () {
        spmStatus = {
            id: "active_status",
            xtype: "syno_displayfield",
            value: _V("ui", "status_unknown"),
            width: 140,
        };
        active = false;
        Ext.Ajax.request({
            url: "/webman/3rdparty/SimplePermissionManager/cgi/status.cgi",
            method: "GET",
            async: false,
            timeout: 60000,
            success: function (response) {
                var data = Ext.decode(response.responseText);
                if (data.active) {
                    spmStatus.value = _V("ui", "status_active");
                    spmStatus.style = {
                        color: "green",
                    };
                    active = true;
                } else {
                    spmStatus.value = _V("ui", "status_inactive");
                }
            },
            failure: function (response) {
                window.alert("Fetch Status Failed.");
            },
        });
        return new SYNO.ux.FieldSet({
            title: _V("ui", "status"),
            collapsible: false,
            items: [
                {
                    xtype: "syno_compositefield",
                    hideLabel: true,
                    items: [
                        {
                            xtype: "syno_displayfield",
                            value: _V("ui", "status") + _T("common", "pure_colon"),
                            width: 140,
                        },
                        spmStatus,
                        {
                            id: "active_button",
                            xtype: "syno_button",
                            btnStyle: "blue",
                            text: _V("ui", "active"),
                            hidden: active,
                            handler: this.onActive.bind(this),
                        },
                    ],
                },
            ],
        });
    },
    // Create the display of config
    createDisplayConfig: function () {
        config = {};
        Ext.Ajax.request({
            url: "/webman/3rdparty/SimplePermissionManager/cgi/config.cgi",
            method: "GET",
            async: false,
            timeout: 60000,
            success: function (response) {
                config = Ext.decode(response.responseText);
            },
            failure: function (response) {
                window.alert("Fetch Config Failed.");
            },
        });
        return new SYNO.ux.FieldSet({
            title: _V("ui", "configure"),
            collapsible: false,
            items: [
                {
                    xtype: "syno_compositefield",
                    hideLabel: true,
                    items: [
                        {
                            xtype: "syno_displayfield",
                            value: _V("ui", "trust_signature") + _T("common", "pure_colon"),
                            width: 140,
                        },
                        {
                            xtype: "syno_checkbox",
                            boxLabel: " ",
                            id: "valid_signature",
                            checked: config.trustSignature,
                            readOnly: true,
                        },
                    ],
                },
            ],
        });
    },
    fetchSynoConfirmPWToken: function (callback) {
        this.sendWebAPI({
            api: "SYNO.Core.User.PasswordConfirm",
            method: "auth",
            version: 2,
            params: {
                password: Ext.getCmp("confirm_password").getValue(),
            },
            callback: function (success, response) {
                if (!success) {
                    window.alert("invalid admin password");
                    return;
                }

                if (
                    response.SynoConfirmPWToken === null ||
                    (typeof response.SynoConfirmPWToken === "string" &&
                        response.SynoConfirmPWToken.trim() === "")
                ) {
                    console.log("empty SynoConfirmPWToken");
                    return;
                }

                callback(response.SynoConfirmPWToken);
            },
            scope: this,
        });
    },
    sendCreateSchedulerTaskWebAPI: function (token) {
        this.sendWebAPI({
            api: "SYNO.Core.EventScheduler.Root",
            method: "create",
            version: 1,
            params: {
                task_name: "Active Simple Permission Manager",
                owner: { 0: "root" },
                event: "bootup",
                enable: true,
                depend_on_task: "",
                notify_enable: false,
                notify_mail: "",
                notify_if_error: false,
                operation_type: "script",
                operation:
                    "/var/packages/SimplePermissionManager/target/bin/spm-update",
                SynoConfirmPWToken: token,
            },
            callback: function (success, message) {
                if (!success) {
                    console.log("error create EventScheduler task");
                    return;
                }

                this.runSchedulerTask();
            },
            scope: this,
        });
    },
    createAndRunSchedulerTask: function () {
        this.fetchSynoConfirmPWToken(
            this.sendCreateSchedulerTaskWebAPI.bind(this)
        );
    },
    sendRunSchedulerTaskWebAPI: function (token) {
        this.sendWebAPI({
            api: "SYNO.Core.EventScheduler",
            method: "run",
            version: 1,
            params: {
                task_name: "Active Simple Permission Manager",
                SynoConfirmPWToken: token,
            },
            callback: function (success, message, data) {
                if (!success) {
                    console.log("error run EventScheduler task");
                    return;
                }
                this.deleteSchedulerTask();
            },
            scope: this,
        });
    },
    runSchedulerTask: function () {
        this.fetchSynoConfirmPWToken(
            this.sendRunSchedulerTaskWebAPI.bind(this)
        );
    },
    sendDeleteSchedulerTaskWebAPI: function (token) {
        this.sendWebAPI({
            api: "SYNO.Core.EventScheduler",
            method: "delete",
            version: 1,
            params: {
                task_name: "Active Simple Permission Manager",
                SynoConfirmPWToken: token,
            },
            callback: function (success, message, data) {
                if (!success) {
                    console.log("error delete EventScheduler task");
                    return;
                }

                Ext.getCmp("confirm_password_dialog").close();
                Ext.getCmp("active_button").hide();
                Ext.getCmp("active_status").setValue("Active");
                const element = document.getElementById("active_status");
                element.style.color = "green";
            },
            scope: this,
        });
    },
    deleteSchedulerTask: function () {
        this.fetchSynoConfirmPWToken(
            this.sendDeleteSchedulerTaskWebAPI.bind(this)
        );
    },
    // Call Active on click
    onActive: function () {
        var window = new SYNO.SDS.ModalWindow({
            id: "confirm_password_dialog",
            title: _T("common", "enter_password_to_continue"),
            width: 500,
            height: 200,
            resizable: !1,
            layout: "fit",
            buttons: [
                {
                    xtype: "syno_button",
                    text: _T("common", "alt_cancel"),
                    scope: this,
                    handler: function () {
                        Ext.getCmp("confirm_password_dialog").close();
                    },
                },
                {
                    xtype: "syno_button",
                    text: _T("common", "submit"),
                    btnStyle: "blue",
                    scope: this,
                    handler: this.createAndRunSchedulerTask.bind(this),
                },
            ],
            items: [
                {
                    xtype: "syno_formpanel",
                    id: "password_form_panel",
                    bodyStyle: "padding: 0",
                    items: [
                        {
                            xtype: "syno_displayfield",
                            value: String.format(
                                _T("common", "enter_user_password")
                            ),
                        },
                        {
                            xtype: "syno_textfield",
                            fieldLabel: _T("common", "password"),
                            textType: "password",
                            id: "confirm_password",
                        },
                    ],
                },
            ],
        });
        window.open();
    },
    // Grid search
    createFilter: function (gridStore) {
        var searchField = new SYNO.ux.TextFilter({
            emptyText: "Search",
            store: gridStore,
            pageSize: 5,
            width: 300,
        });

        var toolbar = new SYNO.ux.Toolbar({
            items: [searchField],
        });

        return toolbar;
    },
    // Create the display of packages
    createPackages: function () {
        return new SYNO.ux.FieldSet({
            title: _V("ui", "package"),
            collapsible: true,
            autoHeight: true,
            items: [
                {
                    xtype: "syno_compositefield",
                    hideLabel: true,
                    items: [this.createPackageGrid()],
                },
            ],
        });
    },
    // Create JSON Store grid calling list packages API
    createPackageGrid: function () {
        var localUrl =
            "/webman/3rdparty/SimplePermissionManager/cgi/list-packages.cgi";

        var gridStore = new SYNO.API.JsonStore({
            autoDestroy: true,
            url: localUrl,
            restful: true,
            root: "result",
            idProperty: "package",
            fields: [
                {
                    name: "enabled",
                    type: "boolean",
                },
                // {
                //     name: "id",
                //     type: "int",
                // },
                {
                    name: "package",
                    type: "string",
                },
                {
                    name: "version",
                    type: "string",
                },
                {
                    name: "description",
                    type: "string",
                },
            ],
        });

        var paging = new SYNO.ux.PagingToolbar({
            store: gridStore,
            displayInfo: true,
            pageSize: 10,
            refreshText: "Reload",
        });

        // var checkboxSelModel = new Ext.grid.CheckboxSelectionModel({
        //     width: 30,
        //     dataIndex: "enabled",
        //     renderer: function(b, c, a) {
        //         return '<div class="x-grid3-row-checker">&#160;</div>'
        //     },
        //     listeners: {
        //         selectionchange: {
        //             fn: function (e) {
        //                 console.log(e);
        //             },
        //             scope: this,
        //         },
        //     },
        // });
        var enableColumn = new SYNO.ux.EnableColumn({
            // header: _T("common", "enabled"),
            dataIndex: "enabled",
            width: 20,
            align: "center",
            bindRowClick: true,
            commitChanges: true,
            disableSelectAll: true,
            enableFastSelectAll: false,
            toggleRec: function(t) {
                var v = t.get(this.dataIndex);
                t.set(this.dataIndex, !v);
                t.json.enabled = !v;

                Ext.Ajax.request({
                    url: "/webman/3rdparty/SimplePermissionManager/cgi/update-package.cgi",
                    method: "POST",
                    timeout: 60000,
                    jsonData: {
                        package: t.json.package,
                        enabled: !v,
                    },
                    success: function (response) {
                        var result = response.responseText;
                        console.log("update-package.cgi response :\n" + result);
                    },
                    failure: function (response) {
                        window.alert("Request Failed.");
                    },
                });
            },
        });

        var c = {
            store: gridStore,
            plugins: [enableColumn],
            colModel: new Ext.grid.ColumnModel({
                defaults: {
                    sortable: true,
                    menuDisabled: true,
                    width: 100,
                    height: 20,
                },
                columns: [
                    // checkboxSelModel,
                    enableColumn,
                    // {
                    //     header: "Enabled",
                    //     width: 30,
                    //     dataIndex: "enabled",
                    // },
                    // {
                    //     header: "Index",
                    //     width: 30,
                    //     dataIndex: "id",
                    // },
                    {
                        header: "Package Name",
                        width: 150,
                        dataIndex: "package",
                    },
                    {
                        header: "Version",
                        width: 100,
                        dataIndex: "version",
                    },
                ],
            }),
            selModel: new Ext.grid.RowSelectionModel(),
            viewConfig: {
                forceFit: true,
                onLoad: Ext.emptyFn,
                listeners: {
                    beforerefresh: function (f) {
                        f.scrollTop = f.scroller.dom.scrollTop;
                    },
                    refresh: function (f) {
                        f.scroller.dom.scrollTop = f.scrollTop;
                    },
                },
            },
            columnLines: true,
            frame: false,
            bbar: paging,
            height: 385,
            cls: "resource-monitor-performance",
            listeners: {
                scope: this,
                render: function (grid) {
                    grid.getStore().load({
                        params: {
                            offset: 0,
                            limit: 10,
                        },
                    });
                },
            },
        };

        return new SYNO.ux.GridPanel(c);
    },
    // Create the display of users
    createUsers: function () {
        return new SYNO.ux.FieldSet({
            title: _V("ui", "user"),
            collapsible: true,
            autoHeight: true,
            items: [
                {
                    xtype: "syno_compositefield",
                    hideLabel: true,
                    items: [this.createUserGrid()],
                },
            ],
        });
    },
    // Create JSON Store grid calling list user API
    createUserGrid: function () {
        var localUrl =
            "/webman/3rdparty/SimplePermissionManager/cgi/list-users.cgi";

        var gridStore = new SYNO.API.JsonStore({
            autoDestroy: true,
            url: localUrl,
            restful: true,
            root: "result",
            idProperty: "name",
            fields: [
                {
                    name: "enabled",
                    type: "boolean",
                },
                {
                    name: "name",
                    type: "string",
                },
                {
                    name: "uid",
                    type: "int",
                },
                {
                    name: "gid",
                    type: "int",
                },
            ],
        });

        var paging = new SYNO.ux.PagingToolbar({
            store: gridStore,
            displayInfo: true,
            pageSize: 10,
            refreshText: "Reload",
        });

        var enableColumn = new SYNO.ux.EnableColumn({
            dataIndex: "enabled",
            width: 20,
            align: "center",
            bindRowClick: true,
            commitChanges: true,
            disableSelectAll: true,
            enableFastSelectAll: false,
            toggleRec: function(t) {
                var v = t.get(this.dataIndex);
                t.set(this.dataIndex, !v);
                t.json.enabled = !v;

                Ext.Ajax.request({
                    url: "/webman/3rdparty/SimplePermissionManager/cgi/update-user.cgi",
                    method: "POST",
                    timeout: 60000,
                    jsonData: {
                        name: t.json.name,
                        enabled: !v,
                    },
                    success: function (response) {
                        var result = response.responseText;
                        console.log("update-user.cgi response :\n" + result);
                    },
                    failure: function (response) {
                        window.alert("Request Failed.");
                    },
                });
            },
        });

        var c = {
            store: gridStore,
            plugins: [enableColumn],
            colModel: new Ext.grid.ColumnModel({
                defaults: {
                    sortable: true,
                    menuDisabled: true,
                    width: 100,
                    height: 20,
                },
                columns: [
                    enableColumn,
                    {
                        header: "User ID",
                        width: 50,
                        dataIndex: "uid",
                    },
                    {
                        header: "User Name",
                        width: 100,
                        dataIndex: "name",
                    },
                    {
                        header: "Group ID",
                        width: 100,
                        dataIndex: "gid",
                    },
                ],
            }),
            selModel: new Ext.grid.RowSelectionModel(),
            viewConfig: {
                forceFit: true,
                onLoad: Ext.emptyFn,
                listeners: {
                    beforerefresh: function (f) {
                        f.scrollTop = f.scroller.dom.scrollTop;
                    },
                    refresh: function (f) {
                        f.scroller.dom.scrollTop = f.scrollTop;
                    },
                },
            },
            columnLines: true,
            frame: false,
            bbar: paging,
            height: 385,
            cls: "resource-monitor-performance",
            listeners: {
                scope: this,
                render: function (grid) {
                    grid.getStore().load({
                        params: {
                            offset: 0,
                            limit: 10,
                        },
                    });
                },
            },
        };

        return new SYNO.ux.GridPanel(c);
    },
});
