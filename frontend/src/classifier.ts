import { plugin } from "./app";
interface Classifier {
    classify(pluigins: plugin.Plugin[], dist: string): Map<plugin.Plugin, string>;
}

export class ByVendor implements Classifier {
    constructor(
        private vendors: string[], // 需要被分类的插件的厂商列表
        private throuthold: number // 分类数量的阈值，对应厂商的插件数量超过阈值时才进行分类
    ) {}
    classify(plugins: plugin.Plugin[], dist: string): Map<plugin.Plugin, string> {
        let P = plugins.reduce<Map<string, plugin.Plugin[]>>((acc, cur) => {
            if (this.vendors.length === 0 || this.vendors.includes(cur.Vendorname)) {
                if (!acc.get(cur.Vendorname)) {
                    acc.set(cur.Vendorname, []);
                }
                acc.get(cur.Vendorname)!.push(cur);
            }
            return acc;
        }, new Map<string, plugin.Plugin[]>());
        for (let vendor of P.keys()) {
            if (P.get(vendor)!.length < this.throuthold) {
                P.delete(vendor);
            }
        }

        const result = new Map<plugin.Plugin, string>();
        for (let [vendor, ps] of P) {
            for (let p of ps) {
                result.set(p, dist + "/" + vendor);
            }
        }
        return result;
    }
}
