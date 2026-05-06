import sys

sp = []  # Список для истории

def add(a, b):
    return a + b

def subtract(a, b):
    return a - b

def multiply(a, b):
    return a * b

def divide(a, b):
    if b == 0:
        return 'Ошибка: деление на 0'
    return a / b
def calculate(n1, op, n2):
    if op == '+':
        res = add(n1, n2)
    elif op == '-':
        res = subtract(n1, n2)
    elif op == '*':
        res = multiply(n1, n2)
    elif op == '/':
        res = divide(n1, n2)
    else:
        return 'Неизвестная операция'
    return res

def show_history():
    print('\n--- История ---')
    for i, entry in enumerate(sp[-5:], 1):
        print(f'  {i}. {entry}')
    print()

def calc():
    while True:
        try:
            n1 = float(input('Первое число: '))
            n2 = float(input('Второе число: '))
        except ValueError:
            print('Ошибка: введите число!')
            continue

        op = input('Операция (+, -, *, /): ')

        result = calculate(n1, op, n2)
        if isinstance(result, str):
            print(result)
            continue

        entry = f'{n1} {op} {n2} = {result}'
        sp.append(entry)
        print(entry)

        pol = input('Ещё раз? (да/history/нет): ')
        if pol == 'нет':
            break
        elif pol == 'history':
            show_history()
        elif pol != 'да':
            print('Не тот выбор')

calc()
