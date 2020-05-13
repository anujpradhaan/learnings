object ClassTutorial{
    def main(args:Array[String]){
        val rover = new Animal
        rover.setName("Anuj")
        rover.setSound("Anger")

        printf("Name is %s and saying %s\n", rover.getName, rover.getSound)
    }

    class Animal(var name: String, var sound: String){
        this.setName(name)

        val id = Animal.newIdNum

        def getName() : String  = name
        def getSound() : String = sound

        def setName(name: String){
            this.name = name
        }

        def setSound(sound: String){
            this.sound = sound
        }

        override def toString(): String = {
            return "%s with id %d says %s".format(this.name,this.id,this.sound)
        }

        def this(){
           this("No Name","No Sound")     
        }

    }

    object Animal{
        private var idNumber = 0
        private def newIdNum = {idNumber+1}
    }
}
