import sys

pol = input('Начнем(да/нет) ')
sp = []

def poll():
    pol1 = float(input('Температура: '))
    pol2 = input('Единица (C/F): ')
    if pol2=='C':
        x = pol1*1.8+32
        xx = f'{pol1}C = {x}'
        sp.append(xx)
        print(f'{pol1}C = {x}')
    elif pol2=='F':
        f = (pol1 - 32)/1.8
        ff = f'{pol1}F = {f}'
        sp.append(ff)
        print(ff)
    else:
        print('Введите верную букву')

def his():
    print(sp)

def circle():
    while True:
        if pol=='да':
            poll()
            print('1. Продолжить\n2. Закончить\n3. История')
            po = input('Число ')
            if po=='1':
                continue
            elif po=='2':
                sys.exit(1)
            elif po=='3':
                his() # Костыль
            else:
                print('Не то число')
        elif pol=='нет':
            sys.exit(1)
        else:
            print('нет')
circle()
