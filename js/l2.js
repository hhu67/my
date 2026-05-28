const readline = require('readline').createInterface({
  input: process.stdin,
  output: process.stdout
});

readline.question('Число ', (answer) => {
  let pol = Number(answer);
  readline.close();
  let x = 0;
  while(x <= 100) { 
        x = x + 1;
        console.log(pol, '*', x, '=', pol * x);
  }
});
