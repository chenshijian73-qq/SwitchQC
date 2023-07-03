export namespace logic {
	
	export class Qc {
	    ID: number;
	    name: string;
	    filepath: string;
	    content: string;
	    enabled: boolean;
	
	    static createFrom(source: any = {}) {
	        return new Qc(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.ID = source["ID"];
	        this.name = source["name"];
	        this.filepath = source["filepath"];
	        this.content = source["content"];
	        this.enabled = source["enabled"];
	    }
	}

}

