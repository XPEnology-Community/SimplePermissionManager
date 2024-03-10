SPK_NAME = SimplePermissionManager
SPK_VERS = 1.0.0
SPK_REV = 7
SPK_ICON = src/SimplePermissionManager.png
WIZARDS_DIR = src/wizard
DSM_UI_DIR = app
# workaround for creating package icons, while ADMIN_URL and SERVICE_PORT is none
DSM_UI_CONFIG = src/app/config

MAINTAINER = jim3ma

DESCRIPTION = Simple Permission Manager is used to approve some packages' priviledged commands automatically, some drivers packages and other packages need high permission to execute their actions.
DESCRIPTION_CHS = 权限管理器用于自动批准其他套件的特权命令, 一些驱动套件需要特权命令才可正常执行.

DISPLAY_NAME = Simple Permission Manager
DISPLAY_NAME_CHS = 权限管理器

STARTABLE = no

HOMEPAGE = https://github.com/XPEnology-Community/SimplePermissionManager

CONF_DIR = src/conf
SYSTEM_GROUP = http

SERVICE_USER   = auto
SERVICE_SETUP  = src/service-setup.sh

COPY_TARGET = nop

POST_STRIP_TARGET = SimplePermissionManager_extra_install

include ../../mk/spksrc.spk.mk

.PHONY: SimplePermissionManager_extra_install
SimplePermissionManager_extra_install:
	install -m 755 -d $(STAGING_DIR)/app/

	ln -s /var/packages/SimplePermissionManager/target/cgi/ $(STAGING_DIR)/app/cgi
	install -m 644 src/app/SimplePermissionManager.js $(STAGING_DIR)/app/SimplePermissionManager.js
	install -m 644 src/app/config $(STAGING_DIR)/app/config
	install -m 644 src/app/helptoc.conf $(STAGING_DIR)/app/helptoc.conf

	install -m 755 -d $(STAGING_DIR)/app/help
	for language in enu; do \
		install -m 755 -d $(STAGING_DIR)/app/help/$${language}; \
		install -m 644 src/app/help/$${language}/SimplePermissionManager_index.html $(STAGING_DIR)/app/help/$${language}/SimplePermissionManager_index.html; \
	done
	install -m 755 -d $(STAGING_DIR)/app/texts
	for language in chs enu; do \
		install -m 755 -d $(STAGING_DIR)/app/texts/$${language}; \
		install -m 644 src/app/texts/$${language}/strings $(STAGING_DIR)/app/texts/$${language}/strings; \
	done

	install -m 755 -d $(STAGING_DIR)/cgi/
	install -m 755 src/cgi/config.cgi $(STAGING_DIR)/cgi/config.cgi
	install -m 755 src/cgi/list-packages.cgi $(STAGING_DIR)/cgi/list-packages.cgi
	install -m 755 src/cgi/list-users.cgi $(STAGING_DIR)/cgi/list-users.cgi
	install -m 755 src/cgi/status.cgi $(STAGING_DIR)/cgi/status.cgi
	install -m 755 src/cgi/update-package.cgi $(STAGING_DIR)/cgi/update-package.cgi
	install -m 755 src/cgi/update-user.cgi $(STAGING_DIR)/cgi/update-user.cgi

	install -m 755 -d $(STAGING_DIR)/bin/
	install -m 755 src/bin/init-permision.py $(STAGING_DIR)/bin/init-permision.py
	install -m 755 src/bin/spm-exec $(STAGING_DIR)/bin/spm-exec
	install -m 755 src/bin/spm-update $(STAGING_DIR)/bin/spm-update
	install -m 644 src/bin/spm-update.sig $(STAGING_DIR)/bin/spm-update.sig

	install -m 755 -d $(STAGING_DIR)/etc.defaults/
	install -m 644 src/etc.defaults/config.json $(STAGING_DIR)/etc.defaults/config.json
	install -m 644 src/etc.defaults/permissions.json $(STAGING_DIR)/etc.defaults/permissions.json
