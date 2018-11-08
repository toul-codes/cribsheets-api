export class Cribsheet {
    private professor: string;
    private coursenumber: string;
    private description: string;
    private filename: string;

    constructor(professor: string, coursenumber: string, description: string, filename: string){
        this.professor = professor;
        this.coursenumber = coursenumber;
        this.description = description;
        this.filename = filename; 
    }

    public getProfessor(){
        return this.professor;
    }

    public getCourseNumber(){
        return this.coursenumber;
    }

    public getDescription(){
        return this.description;
    }

    public setName(name: string){
        this.professor = name;
    }

    public setDescription(description: string){
        this.description = description;
    }

}
