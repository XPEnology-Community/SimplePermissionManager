SPK_NAME = SimplePermissionManager
SPK_VERS = 0.2
SPK_REV = 1
SPK_ICON = src/SimplePermissionManager.png
WIZARDS_DIR = WIZARD_UIFILES
DSM_UI_DIR = ui

MAINTAINER = jim3ma

DESCRIPTION = Simple Permission Manager
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
	install -m 755 src/cgi/perl.cgi $(STAGING_DIR)/cgi/perl.cgi
	install -m 755 src/cgi/python.cgi $(STAGING_DIR)/cgi/python.cgi
	install -m 755 src/cgi/bash.cgi $(STAGING_DIR)/cgi/bash.cgi
	install -m 755 src/cgi/externalapi.cgi $(STAGING_DIR)/cgi/externalapi.cgi
	install -m 755 src/cgi/storepythonsqlapi.cgi $(STAGING_DIR)/cgi/storepythonsqlapi.cgi
	install -m 755 src/cgi/storepythonsynoapi.cgi $(STAGING_DIR)/cgi/storepythonsynoapi.cgi
	install -m 755 src/cgi/storebashratesapi.cgi $(STAGING_DIR)/cgi/storebashratesapi.cgi

	install -m 755 src/cgi/synopkg-list-packages.cgi $(STAGING_DIR)/cgi/synopkg-list-packages.cgi
	install -m 755 src/cgi/status.cgi $(STAGING_DIR)/cgi/status.cgi
	install -m 755 src/cgi/update-package.cgi $(STAGING_DIR)/cgi/update-package.cgi

	install -m 755 -d $(STAGING_DIR)/bin/
	install -m 755 src/bin/spm-exec $(STAGING_DIR)/bin/spm-exec
