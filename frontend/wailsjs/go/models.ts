export namespace main {
	
	export class Form {
	    name: string;
	    content: string;
	    enabled: boolean;
	
	    static createFrom(source: any = {}) {
	        return new Form(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.name = source["name"];
	        this.content = source["content"];
	        this.enabled = source["enabled"];
	    }
	}

}

export namespace tables {
	
	export class Qc {
	    // Go type: gtime
	    createAt?: any;
	    // Go type: gtime
	    updateAt?: any;
	    // Go type: gtime
	    deleteAt?: any;
	
	    static createFrom(source: any = {}) {
	        return new Qc(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.createAt = this.convertValues(source["createAt"], null);
	        this.updateAt = this.convertValues(source["updateAt"], null);
	        this.deleteAt = this.convertValues(source["deleteAt"], null);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}

}

