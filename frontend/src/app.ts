import { GetConfig, MovePlugin, SaveConfig, ScanPluginDB } from "../wailsjs/go/main/App";
import { config as cfg, main, nfo } from "../wailsjs/go/models";

export namespace plugin {
    export async function Scan(): Promise<Plugin[]> {
        return (await ScanPluginDB()).map((p) => new Plugin(p));
    }
    export class Plugin extends main.Plugin {
        declare readonly Nfo: nfo.Plugin;
        declare readonly PresetPath: string;
        declare readonly Name: string;
        declare readonly FstName: string;
        declare readonly NfoName: string;
        declare readonly Vendorname: string;
        declare readonly Bitsize: number;
        declare readonly Category: string[];
        GetCoverURL(): string {
            return `url(data:${this.CoverMimeType};base64,${this.Cover})`;
        }
        async MoveTo(dist: string) {
            const p = await MovePlugin(this, dist);
            Object.assign(this, p);
        }
    }
}

export namespace config {
    export async function Get(): Promise<cfg.Config> {
        return GetConfig();
    }
    export async function Save(config: cfg.Config) {
        return SaveConfig(config);
    }
}
