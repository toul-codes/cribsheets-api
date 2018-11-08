export class Cribsheet {
    private lastname: string;
    private subject: string;
    private type: string;
    private university: string;
	private year: string;

    constructor(lastName: string, subject: string, type: string, university: string, year: string){
        this.lastname = lastName;
        this.subject = subject;
        this.type = type;
        this.university = university; 
		this.year = year;
    }

    public getLastName(){
        return this.lastname;
    }

    public getSubject(){
        return this.subject;
    }

	public getType(){
		return this.type;
	}

	public getUniversity(){
		return this.university;
    }
    
    public getYear(){
        return this.year;
    }

}
