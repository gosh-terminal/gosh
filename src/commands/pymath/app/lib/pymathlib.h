#include <cstdlib>
#include <fstream>
#include <iostream>
#include <string>
using namespace std;
namespace pymath {
    void algebra() {
        string input1;
        cout << "Welcome to the algebra part of pymath available calculators include logorithm calc(log), and quadratic equation calc(quad)!\nWhich do you want?: ";
        cin >> input1;
        if (input1 == "log") {
            string base, result;
            cout << "What is the base number?: ";
            cin >> base;
            cout << "What is the result?: ";
            cin >> result;
            ofstream outfile;
            outfile.open("app/data/output.txt");
            outfile << base << endl << result << endl;
            system("python3 app/lib/logexe.py < app/data/output.txt");
        } else if (input1 == "quad") {
            string a,b,c;
            cout << "What is a?: ";
            cin >> a;
            cout << "What is b?: ";
            cin >> b;
            cout << "What is c?: ";
            cin >> c;
            ofstream outfile;
            outfile.open("app/data/output.txt");
            outfile << a << endl << b << endl << c << endl;
            system("python3 app/lib/quadexe.py < app/data/output.txt");
        }
    }
    void geometry() {
        string input1;
        cout << "Welcome to the geometry part of pymath available calculators include circumfrence calculator(circ), pythagorean theorem calc(pythag), area calc(area)\nWhich do you want?: ";
        cin >> input1;
        if (input1 == "circ") {
            string circumfrence;
            cout << "What is the diameter?: ";
            cin >> circumfrence;
            ofstream outfile;
            outfile.open("app/data/output.txt");
            outfile << circumfrence << endl;
            system("python3 app/lib/circexe.py < app/data/output.txt");
        } else if (input1 == "pythag") {
            string a,b;
            cout << "What is a?: ";
            cin >> a;
            cout << "What is b?: ";
            cin >> b;
            ofstream outfile;
            outfile.open("app/data/output.txt");
            outfile << a << endl << b << endl;
            system("python3 app/lib/pythagexe.py < app/data/output.txt");
        } else if(input1 == "area") {
            cout << "Welcome to the area calculator from the geometry calculator! available calculators include area of circle(aoc), area of trapezoid(aoz), area of triangle(aotri)\n Which do you want?: ";
            string input2;
            cin >> input2;
            if (input2 == "aoc") {
                string radius;
                cout << "What is the radius?: ";
                cin >> radius;
                ofstream outfile;
                outfile.open("app/data/output.txt");
                outfile << radius << endl;
                system("python3 app/lib/aocexe.py < app/data/output.txt");
            } else if (input2 == "aoz")  {
                string a,b,h;
                cout << "What is a?: ";
                cin >> a;
                cout << "What is b?: ";
                cin >> b;
                cout << "What's the height?: ";
                cin >> h;
                ofstream outfile;
                outfile.open("app/data/output.txt");
                outfile << a << endl << b << endl << h << endl;
                system("python3 app/lib/aozexe.py < app/data/output.txt");
            } else if (input2 == "aotri") {
                string b,h;
                cout << "What is the base?: ";
                cin >> b;
                cout << "What is the height?: ";
                cin >> h;
                ofstream outfile;
                outfile.open("app/data/output.txt");
                outfile << b << endl << h << endl;
                system("python3 app/lib/aotriexe.py < app/data/output.txt");
            }
        }
    }
    void physics() {
        cout << "Welcome to the physics portion of pymath! available calculators include acceleration calc(acc), displacement calculator(disp), speed calc(speed), tangental acceleration calc(tangacc), velocity calc (velocity), tangental speed calc(tanspeed)\nWhich do you want?: ";
        string input;
        cin >> input;
        if (input == "acc") {
            string v1,v2,t1,t2;
            cout << "What is the initial velocity?: ";
            cin >> v1;
            cout << "What is the final velocity?: ";
            cin >> v2;
            cout << "What is the initial time?: ";
            cin >> t1;
            cout << "What is the final time?: ";
            cin >> t2;
            ofstream outfile;
            outfile.open("app/data/output.txt");
            outfile << v1 << endl << v2 << endl << t1 << endl << t2 << endl;
            system("python3 app/lib/accexe.py < app/data/output.txt");
        } else if (input == "disp") {
            string initdis,findis,initime,fintime;
            cout << "What is the initial displacement?: ";
            cin >> initdis;
            cout << "What is the final displacement?: ";
            cin >> findis;
            cout << "What is the initial time?: ";
            cin >> initime;
            cout << "What is the final time?: ";
            cin >> fintime;
            ofstream outfile;
            outfile.open("app/data/output.txt");
            outfile << initdis << endl << findis << endl << initime << endl << fintime << endl;
            system("python3 app/lib/displexe.py < app/data/output.txt");
        } else if (input == "speed") {
            string distance,time;
            cout << "What is the distance?: ";
            cin >> distance;
            cout << "What is the time?: ";
            cin >> time;
            ofstream outfile;
            outfile.open("app/data/output.txt");
            outfile << distance << endl << time << endl;
            system("python3 app/lib/speedexe.py < app/data/output.txt");
        } else if (input == "tangacc") {
            string radius,time,cycles;
            cout << "What is the radius?: ";
            cin >> radius;
            cout << "What is the time?: ";
            cin >> time;
            cout << "How many cycles are there?: ";
            cin >> cycles;
            ofstream outfile;
            outfile.open("app/data/output.txt");
            outfile << radius << endl << time << endl << cycles << endl;
            system("python3 app/lib/tanaccexe.py < app/data/output.txt");
        } else if (input == "velocity") {
            string displacement,speed;
            cout << "What is the displacement?: ";
            cin >> displacement;
            cout << "What is the speed?: ";
            cin >> speed;
            ofstream outfile;
            outfile.open("app/data/output.txt");
            outfile << displacement << endl << speed << endl;
            system("python3 app/lib/velexe.py < app/data/output.txt");
        } else if (input == "tanspeed") {
            string radius,time,cycles;
            cout << "What is the radius?: ";
            cin >> radius;
            cout << "What is the time?: ";
            cin >> time;
            cout << "How many cycles are there?: ";
            cin >> cycles;
            ofstream outfile;
            outfile.open("app/data/output.txt");
            outfile << radius << endl << time << endl << cycles << endl;
            system("python3 app/lib/tanspeedexe.py < app/data/output.txt");
        }
    }
    void chemistry() {
        //WIP
    }
}