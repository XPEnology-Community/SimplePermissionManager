SPK_NAME = SimplePermissionManager
SPK_VERS = 0.2
SPK_REV = 1
SPK_ICON = src/SimplePermissionManager.png
WIZARDS_DIR = WIZARD_UIFILES
DSM_UI_DIR = ui

MAINTAINER = jim3ma

DESCRIPTION = Simple Permission Manager is used to approve priviledged commands automatically, some drivers packages and other packages need high permission to execute their actions
STARTABLE = no
DISPLAY_NAME = Simple Permission Manager

HOMEPAGE = https://github.com/jim3ma/todo

ADMIN_URL = /SimplePermissionManager/

CONF_DIR = src/conf
SYSTEM_GROUP = http

SERVICE_USER   = auto
SERVICE_SETUP  = src/service-setup.sh

COPY_TARGET = nop

POST_STRIP_TARGET = SimplePermissionManager_extra_install

include ../../mk/spksrc.spk.mk

.PHONY: SimplePermissionManager_extra_install
SimplePermissionManager_extra_install:
	install -m 755 -d $(STAGING_DIR)/ui/

	ln -s /var/packages/SimplePermissionManager/target/cgi/ $(STAGING_DIR)/ui/cgi
	install -m 644 src/ui/SimplePermissionManager.js $(STAGING_DIR)/ui/SimplePermissionManager.js
	install -m 644 src/ui/config $(STAGING_DIR)/ui/config
	install -m 644 src/ui/helptoc.conf $(STAGING_DIR)/ui/helptoc.conf

	install -m 755 -d $(STAGING_DIR)/ui/help
	for language in enu; do \
		install -m 755 -d $(STAGING_DIR)/ui/help/$${language}; \
		install -m 644 src/ui/help/$${language}/SimplePermissionManager_index.html $(STAGING_DIR)/ui/help/$${language}/SimplePermissionManager_index.html; \
	done
	install -m 755 -d $(STAGING_DIR)/ui/texts
	for language in enu; do \
		install -m 755 -d $(STAGING_DIR)/ui/texts/$${language}; \
		install -m 644 src/ui/texts/$${language}/strings $(STAGING_DIR)/ui/texts/$${language}/strings; \
	done

	install -m 755 -d $(STAGING_DIR)/cgi/
	install -m 755 src/cgi/config.cgi $(STAGING_DIR)/cgi/config.cgi
	install -m 755 src/cgi/list-packages.cgi $(STAGING_DIR)/cgi/list-packages.cgi
	install -m 755 src/cgi/list-users.cgi $(STAGING_DIR)/cgi/list-users.cgi
	install -m 755 src/cgi/status.cgi $(STAGING_DIR)/cgi/status.cgi
	install -m 755 src/cgi/update-package.cgi $(STAGING_DIR)/cgi/update-package.cgi
	install -m 755 src/cgi/update-user.cgi $(STAGING_DIR)/cgi/update-user.cgi

	install -m 755 -d $(STAGING_DIR)/bin/
	install -m 755 src/bin/spm-exec $(STAGING_DIR)/bin/spm-exec
	install -m 755 src/bin/spm-update $(STAGING_DIR)/bin/spm-update
	install -m 644 src/bin/spm-update.sig $(STAGING_DIR)/bin/spm-update.sig

	install -m 755 -d $(STAGING_DIR)/etc.defaults/
	install -m 644 src/etc.defaults/config.json $(STAGING_DIR)/etc.defaults/config.json
	install -m 644 src/etc.defaults/permissions.json $(STAGING_DIR)/etc.defaults/permissions.json
