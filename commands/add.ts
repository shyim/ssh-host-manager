import chalk from "chalk";
import Host from "../struct/Host";
import Helper from "../helper";
import * as fs from "fs";

const parseUrl = require('parse-url');

export default class AddShortcut {
    static execute(argv) {
        if (argv.name === undefined || argv.uri === undefined) {
            console.log(chalk.red('Usage: sshm add [Name] [Connection URI]'));
            return;
        }

        AddShortcut.fixDefaultConfig();

        let uriProperties = parseUrl(argv.uri);

        if (uriProperties.resource === null) {
            console.log(chalk.red('Invalid connection uri'));
            return;
        }

        uriProperties.port = uriProperties.port || 22;

        let host = new Host({
            name: argv.name,
            hostname: uriProperties.resource,
            user: uriProperties.user || process.env['USER'],
            port: uriProperties.port,
            identityfile: argv.identifyfile || null
        });

        let currentConfig = Helper.parseSSHConfig();

        currentConfig.forEach((host) => {
            if (host.name.toLowerCase() === argv.name.toLowerCase()) {
                console.log(chalk.red(`A entry with name ${argv.name} already exists.`));
                process.exit(0);
            }
        });

        currentConfig.push(host);

        Helper.writeSSHConfig(currentConfig);

        console.log(chalk.green(`Successfully added ${host.name}, Type ssh ${host.name} to connect to this server`));
    }

    static fixDefaultConfig() {
        let sshConfigPath = process.env['HOME'] + '/.ssh/config';

        if (fs.existsSync(sshConfigPath)) {
            let content = fs.readFileSync(sshConfigPath).toString();
            if (content.indexOf('manager_hosts') === -1) {
                content += '\nInclude manager_hosts';
                fs.writeFileSync(sshConfigPath, content, {mode: 600});
            }
        } else {
            if (!fs.existsSync(process.env['HOME'] + '/.ssh')) {
                fs.mkdirSync(process.env['HOME'] + '/.ssh');
            }
            fs.writeFileSync(sshConfigPath, 'Include manager_hosts', {mode: 600});
        }
    }
}