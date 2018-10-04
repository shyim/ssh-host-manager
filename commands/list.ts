import Helper from "../helper";
import chalk from "chalk";
import * as CliTable from "cli-table3";
import Host from "../struct/Host";

export default class List {
    static execute() {
        let config = Helper.parseSSHConfig();

        console.log(chalk.green('Entries:'));

        let table = new CliTable({
            head: ['Name', 'Host', 'Port', 'User', 'Identity File'],
            colWidths: [20, 20]
        });

        config.forEach((row) => {
            table.push([row.name, row.hostname, row.port.toString(), row.user, row.identityfile]);
        });

        console.log(table.toString());
    }
}
