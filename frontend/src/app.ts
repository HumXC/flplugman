import {
    GetConfig,
    MovePlugin,
    SaveConfig,
    ScanPluginDB,
    OpenDirectoryDialog,
    OpenFileDialog,
} from "../wailsjs/go/main/App";
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
            if (!this.Cover) return "";
            return `url(${this.GetCoverData()})`;
        }
        GetCoverData(): string {
            if (!this.Cover) return "";
            return `data:${this.CoverMimeType};base64,${this.Cover}`;
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

export namespace utils {
    export async function ChooseDir(title: string = "", defaultDir: string = ""): Promise<string> {
        return await OpenDirectoryDialog(title, defaultDir);
    }
    // await utils.ChooseFile("打开文件", "", [
    //     {
    //         Pattern: "*.*",
    //         DisplayName: "All Files (*.*)",
    //     },
    // ]);
    export async function ChooseFile(
        title: string = "",
        defaultDir: string = "",
        filter: main.FileFilter[] = []
    ): Promise<string> {
        return await OpenFileDialog(title, defaultDir, filter);
    }
}
