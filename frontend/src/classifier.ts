import { main } from "../wailsjs/go/models";
interface Classifier {
    classify(pluigins: main.Plugin[], dist: string): Map<main.Plugin, string>;
}

class ByVendor implements Classifier {
    constructor(
        private vendors: string[], // 需要被分类的插件的厂商列表
        private throuthold: number // 分类数量的阈值，对应厂商的插件数量超过阈值时才进行分类
    ) {}
    classify(plugins: main.Plugin[], dist: string): Map<main.Plugin, string> {
        let P = plugins.reduce<Map<string, main.Plugin[]>>((acc, cur) => {
            if (this.vendors.includes(cur.Vendorname)) {
                if (!acc.get(cur.Vendorname)) {
                    acc.set(cur.Vendorname, []);
                }
                acc.get(cur.Vendorname)!.push(cur);
            }
            return acc;
        }, new Map<string, main.Plugin[]>());
        for (let vendor of P.keys()) {
            if (P.get(vendor)!.length < this.throuthold) {
                P.delete(vendor);
            }
        }
        if (this.vendors.length !== 0) {
            const pp = P;
            P = new Map<string, main.Plugin[]>();
            for (let vendor of this.vendors) {
                if (pp.get(vendor)) {
                    P.set(vendor, pp.get(vendor)!);
                }
            }
        }
        const result = new Map<main.Plugin, string>();
        for (let [vendor, ps] of P) {
            for (let p of ps) {
                result.set(p, dist + "/" + vendor);
            }
        }
        return result;
    }
}
