let x = {
    name: "Алеша",
    age: 25
}
x.city = "ДА";
console.log(x.age);
console.log(x.name);
console.log(x.city);
let al = ["яблоко", "человек", "бананы"];
console.log(al[1]);
if(al[1] == "человек") {
    console.log("сергей");
}
let i = 0;
while(1) {
    console.log("сергей");
    i++;
    if(i < 10) {
        continue;
    }
    else {
        break;
    }
}