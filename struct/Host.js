"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
class Host {
    constructor(config) {
        Object.keys(config).forEach((key) => {
            this[key] = config[key];
        });
    }
}
exports.default = Host;
