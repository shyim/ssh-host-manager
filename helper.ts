import * as fs from "fs";
import Host from "./struct/Host";
import chalk from "chalk";
import {exec} from "child_process";


export default class Helper {
    static getPathToConfig() {
        return process.env['HOME'] + '/.ssh/manager_hosts';
    }

    static parseSSHConfig(): Host[] {
        let path = Helper.getPathToConfig();
        if (!fs.existsSync(path)) {
            return [];
        }

        let fileContent = fs.readFileSync(path).toString();
        let fileRows = fileContent.split("\n");
        let config = [];
        let hostInfo = {};

        fileRows.forEach((row: string) => {
            if (row.toLowerCase().substr(0, 4) === 'host') {
                if (Object.keys(hostInfo).length > 0) {
                    let host = new Host(hostInfo);
                    config.push(host);
                }
                hostInfo = {};
                hostInfo['name'] = row.substr(4).trim();
            } else if (row.substr(0, 1) === ' ') {
                let property = row.trim().split(' ');
                hostInfo[property[0]] = property[1];
            }
        });

        if (Object.keys(hostInfo).length > 0) {
            let host = new Host(hostInfo);
            config.push(host);
        }

        return config;
    }

    static writeSSHConfig(hosts: Host[]) {
        let content = '';

        hosts.forEach((host) => {
            content += `Host ${host.name}\n`;

            Object.keys(host).forEach((key : string) => {
                if (key !== 'name' && host[key]) {
                    content += `    ${key} ${host[key]}\n`;
                }
            });
        });

        fs.writeFileSync(Helper.getPathToConfig(), content, {mode: 600});
    }
}