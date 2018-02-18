"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
const chalk_1 = require("chalk");
const Host_1 = require("../struct/Host");
const helper_1 = require("../helper");
const fs = require("fs");
const parseUrl = require('parse-url');
class AddShortcut {
    static execute(argv) {
        if (argv.name === undefined || argv.uri === undefined) {
            console.log(chalk_1.default.red('Usage: sshm add [Name] [Connection URI]'));
            return;
        }
        AddShortcut.fixDefaultConfig();
        let uriProperties = parseUrl(argv.uri);
        if (uriProperties.resource === null) {
            console.log(chalk_1.default.red('Invalid connection uri'));
            return;
        }
        uriProperties.port = uriProperties.port || 22;
        let host = new Host_1.default({
            name: argv.name,
            hostname: uriProperties.resource,
            user: uriProperties.user || process.env['USER'],
            port: uriProperties.port,
            identityfile: argv.identifyfile || null
        });
        let currentConfig = helper_1.default.parseSSHConfig();
        currentConfig.forEach((host) => {
            if (host.name.toLowerCase() === argv.name.toLowerCase()) {
                console.log(chalk_1.default.red(`A entry with name ${argv.name} already exists.`));
                process.exit(0);
            }
        });
        currentConfig.push(host);
        helper_1.default.writeSSHConfig(currentConfig);
        console.log(chalk_1.default.green(`Successfully added ${host.name}, Type ssh ${host.name} to connect to this server`));
    }
    static fixDefaultConfig() {
        let sshConfigPath = process.env['HOME'] + '/.ssh/config';
        if (fs.existsSync(sshConfigPath)) {
            let content = fs.readFileSync(sshConfigPath).toString();
            if (content.indexOf('manager_hosts') === -1) {
                content += '\nInclude manager_hosts';
                fs.writeFileSync(sshConfigPath, content, { mode: 600 });
            }
        }
        else {
            if (!fs.existsSync(process.env['HOME'] + '/.ssh')) {
                fs.mkdirSync(process.env['HOME'] + '/.ssh');
            }
            fs.writeFileSync(sshConfigPath, 'Include manager_hosts', { mode: 600 });
        }
    }
}
exports.default = AddShortcut;
