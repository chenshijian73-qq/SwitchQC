export namespace logic {
	
	export class Qc {
	    id: number;
	    name: string;
	    filepath: string;
	    content: string;
	    enabled: boolean;
	
	    static createFrom(source: any = {}) {
	        return new Qc(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.name = source["name"];
	        this.filepath = source["filepath"];
	        this.content = source["content"];
	        this.enabled = source["enabled"];
	    }
	}

}

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

