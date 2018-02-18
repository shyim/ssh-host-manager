export default class Host {
    constructor(config : any) {
        Object.keys(config).forEach((key: string) => {
            this[key] = config[key];
        })
    }

    public name : string;
    public hostname : string;
    public port : number;
    public user : string;
    public identityfile: string;
}