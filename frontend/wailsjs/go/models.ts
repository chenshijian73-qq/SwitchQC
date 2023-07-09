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
	    CreateAt?: any;
	    // Go type: gtime
	    UpdateAt?: any;
	    // Go type: gtime
	    DeleteAt?: any;
	
	    static createFrom(source: any = {}) {
	        return new Qc(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.CreateAt = this.convertValues(source["CreateAt"], null);
	        this.UpdateAt = this.convertValues(source["UpdateAt"], null);
	        this.DeleteAt = this.convertValues(source["DeleteAt"], null);
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

