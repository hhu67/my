import sys

lis=[]

while True:
    print('Введите 1, чтобы добавить продукт\n\nВведите стоп если хотите закончить')
    pol=input('Ведите: ')
    if pol=='1':
        add_list=input('Введите предмет ')
        if f'{add_list}' in lis:
            print(f'{add_list} уже есть в списке')
        else:
            add=lis.append(add_list)
    elif pol=='стоп':
        print(f'Ваш список покупок {lis}')
        sys.exit()
    else:
        print('Введите верный')