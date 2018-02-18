"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
const helper_1 = require("../helper");
const chalk_1 = require("chalk");
const CliTable2 = require("cli-table2");
class List {
    static execute() {
        let config = helper_1.default.parseSSHConfig();
        console.log(chalk_1.default.green('Entries:'));
        let table = new CliTable2({
            head: ['Name', 'Host', 'Port', 'User'],
            colWidths: [20, 20]
        });
        config.forEach((row) => {
            table.push([row.name, row.hostname, row.port.toString(), row.user]);
        });
        console.log(table.toString());
    }
}
exports.default = List;
