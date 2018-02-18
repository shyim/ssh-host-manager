"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
const fs = require("fs");
const Host_1 = require("./struct/Host");
class Helper {
    static getPathToConfig() {
        return process.env['HOME'] + '/.ssh/manager_hosts';
    }
    static parseSSHConfig() {
        let path = Helper.getPathToConfig();
        if (!fs.existsSync(path)) {
            return [];
        }
        let fileContent = fs.readFileSync(path).toString();
        let fileRows = fileContent.split("\n");
        let config = [];
        let hostInfo = {};
        fileRows.forEach((row) => {
            if (row.toLowerCase().substr(0, 4) === 'host') {
                if (Object.keys(hostInfo).length > 0) {
                    let host = new Host_1.default(hostInfo);
                    config.push(host);
                }
                hostInfo = {};
                hostInfo['name'] = row.substr(4).trim();
            }
            else if (row.substr(0, 1) === ' ') {
                let property = row.trim().split(' ');
                hostInfo[property[0]] = property[1];
            }
        });
        if (Object.keys(hostInfo).length > 0) {
            let host = new Host_1.default(hostInfo);
            config.push(host);
        }
        return config;
    }
    static writeSSHConfig(hosts) {
        let content = '';
        hosts.forEach((host) => {
            content += `Host ${host.name}
   hostname ${host.hostname}
   user ${host.user}
   port ${host.port}`;
            if (host.identityfile) {
                content += "\n" + '   identityfile ' + host.identityfile;
            }
            content += "\n";
        });
        fs.writeFileSync(Helper.getPathToConfig(), content, { mode: 600 });
    }
}
exports.default = Helper;
