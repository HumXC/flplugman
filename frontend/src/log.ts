import * as l from "../wailsjs/go/log/JSLogger";

export namespace logger {
    export function Info(...message: string[]) {
        l.Info(message.join(" "));
    }
    export function Warn(...message: string[]) {
        l.Warn(message.join(" "));
    }
    export function Error(...message: string[]) {
        l.Error(message.join(" "));
    }
    export function Debug(...message: string[]) {
        l.Debug(message.join(" "));
    }
    export function Fatal(...message: string[]) {
        l.Fatal(message.join(" "));
    }
}
