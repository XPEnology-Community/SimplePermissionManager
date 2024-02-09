SPK_NAME = SimplePermissionManager
SPK_VERS = 0.1
SPK_REV = 1
SPK_ICON = src/simplepermissionmanager.png
WIZARDS_DIR = WIZARD_UIFILES
DSM_UI_DIR = app

MAINTAINER = Jim Ma

DESCRIPTION = Simple Permission Manager
STARTABLE = no
DISPLAY_NAME = Simple Permission Manager

HOMEPAGE = https://github.com/todo

ADMIN_URL = /simplepermissionmanager/

CONF_DIR = src/conf
SYSTEM_GROUP = http

SERVICE_USER   = auto
SERVICE_SETUP  = src/service-setup.sh

COPY_TARGET = nop

POST_STRIP_TARGET = simplepermissionmanager_extra_install

include ../../mk/spksrc.spk.mk

.PHONY: simplepermissionmanager_extra_install
simplepermissionmanager_extra_install:
	install -m 755 -d $(STAGING_DIR)/app/
	install -m 755 src/app/perl.cgi $(STAGING_DIR)/app/perl.cgi
	install -m 755 src/app/python.cgi $(STAGING_DIR)/app/python.cgi
	install -m 755 src/app/bash.cgi $(STAGING_DIR)/app/bash.cgi
	install -m 755 src/app/externalapi.cgi $(STAGING_DIR)/app/externalapi.cgi
	install -m 755 src/app/storepythonsqlapi.cgi $(STAGING_DIR)/app/storepythonsqlapi.cgi
	install -m 644 src/app/config $(STAGING_DIR)/app/config
	install -m 644 src/app/simplepermissionmanager.js $(STAGING_DIR)/app/simplepermissionmanager.js
	install -m 644 src/app/helptoc.conf $(STAGING_DIR)/app/helptoc.conf
	install -m 755 src/app/createsqlitedata.sql $(STAGING_DIR)/app/createsqlitedata.sql
	install -m 755 src/app/storepythonsynoapi.cgi $(STAGING_DIR)/app/storepythonsynoapi.cgi
	install -m 755 src/app/storebashratesapi.cgi $(STAGING_DIR)/app/storebashratesapi.cgi
	install -m 755 -d $(STAGING_DIR)/app/help
	for language in enu; do \
		install -m 755 -d $(STAGING_DIR)/app/help/$${language}; \
		install -m 644 src/app/help/$${language}/simplepermissionmanager_index.html $(STAGING_DIR)/app/help/$${language}/simplepermissionmanager_index.html; \
	done
	install -m 755 -d $(STAGING_DIR)/app/texts
	for language in enu; do \
		install -m 755 -d $(STAGING_DIR)/app/texts/$${language}; \
		install -m 644 src/app/texts/$${language}/strings $(STAGING_DIR)/app/texts/$${language}/strings; \
	done
