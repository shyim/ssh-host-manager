"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
const chalk_1 = require("chalk");
const helper_1 = require("../helper");
class RemoveShortcut {
    static execute(argv) {
        if (argv.name === undefined) {
            console.log(chalk_1.default.red('Usage: sshm remove [Name]'));
            return;
        }
        let hosts = helper_1.default.parseSSHConfig();
        hosts.forEach((host, index) => {
            if (host.name.toLowerCase() === argv.name.toLowerCase()) {
                hosts.splice(index, 1);
                helper_1.default.writeSSHConfig(hosts);
                console.log(chalk_1.default.green(`Entry with name ${host.name} has been removed!`));
                process.exit(0);
            }
        });
        console.log(chalk_1.default.red(`A entry with name ${argv.name} could not found.`));
    }
}
exports.default = RemoveShortcut;
