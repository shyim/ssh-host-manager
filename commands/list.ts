import Helper from "../helper";
import chalk from "chalk";
import * as CliTable2 from "cli-table2";
import Host from "../struct/Host";

export default class List {
    static execute() {
        let config = Helper.parseSSHConfig();

        console.log(chalk.green('Entries:'));

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