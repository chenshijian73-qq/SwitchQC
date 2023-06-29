export namespace logic {
	
	export class QuickCmdFile {
	    name: string;
	    content: string;
	
	    static createFrom(source: any = {}) {
	        return new QuickCmdFile(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.name = source["name"];
	        this.content = source["content"];
	    }
	}

}

