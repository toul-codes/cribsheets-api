export class Cribsheet {
    private name: string;
    private university: string;
    private professor: string;
	private subject: string;
	private type: string;
	private description: string

    constructor(name: string, university: string, professor: string, subject: string, type: string, description: string){
        this.name = name;
        this.university = university;
        this.professor = professor;
		this.subject = subject;
		this.type = type;
		this.description = description;
    }

    public getName(){
        return this.name;
    }

    public getUniversity(){
        return this.university;
    }

    public getProfessor(){
        return this.professor;
    }

	public getSubject(){
		return this.subject;
	}

	public getType(){
		return this.type;
	}

	public getDescription(){
		return this.description;
	}

    public setName(name: string){
        this.name = name;
    }

    public setDescription(description: string){
        this.description = description;
    }
}
