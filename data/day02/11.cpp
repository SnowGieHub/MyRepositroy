#include <iostream>
#include<string>
using namespace std;



class Game{

public:
        virtual void show() = 0;

};

class Hero:public Game{

public:
    Hero(string name)
    {
        this->name = name;
    }

    void show()
    {
    
        cout<<"Hero show...."<<endl;      
    }

    
string  name;

};


class Star:public Game{

    public :
    Star(string color ,int x,int y)
    {
        this->color = color;
        this->x = x;
        this->y = y;
        
    }
    void show()
    {
    
        cout<<"Star show...."<<endl;
    }
    


    string color;
    int x;
    int y;

};

class Stone :public Game{


    public:

        Stone(int x,int y){
            this->x = x;
            this->y = y;
        
        }

    void show()
    {
        cout<<"stone show..."<<endl;
    }


    int x;
    int y;

};

class Test{

    public:

       Game *  Creat(string type)
       {
            if(type == "hero")
            {
                return  new Hero("亚瑟");
                

            }
            else if(type == "star")
            {
                return  new Star("yellow",10,20);

            }
            else if(type == "stone")
            {

                return new Stone(10,70);
            }

       }
    

};



int main(int argc, char *argv[])
{

    Test *test = new Test;

    Game * hero = test->Creat("hero");
    hero->show();

    Game * star = test->Creat("star");
    star->show();


    Game * stone = test->Creat("stone");
    stone->show();
	return 0;
}
