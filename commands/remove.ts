import chalk from "chalk";
import Helper from "../helper";

export default class RemoveShortcut {
    static execute(argv) {
        if (argv.name === undefined) {
            console.log(chalk.red('Usage: sshm remove [Name]'));
            return;
        }

        let hosts = Helper.parseSSHConfig();

        hosts.forEach((host, index) => {
            if (host.name.toLowerCase() === argv.name.toLowerCase()) {
                hosts.splice(index, 1);
                Helper.writeSSHConfig(hosts);
                console.log(chalk.green(`Entry with name ${host.name} has been removed!`));
                process.exit(0);
            }
        });

        console.log(chalk.red(`A entry with name ${argv.name} could not found.`))
    }
}