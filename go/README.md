strings:
    .ToLower(x) - для одного нижнего регистра
    .HasSuffix(x, "aaa") - для проверки оканчивается ли x на ааа(true/false)
    .HasPrefix(x, "aaa") - для проверки начинается ли x на aaa(true/false)
    .Contains(x, "aaa") - для проверки существует ли вообще aaa в x(true/false)
    .ContainsAny(x, "abc") - для проверки есть ли в x буквы a b c
    .Count(x, "aaa") - для продсчета сколько в x aaa
    .EqualFold(x, x2) - сравнивает переменные без учета их регистра 
slice:
    x := []int{} - создание среза
    x := make([]int, a, b) - создание среза с параметром, a - длинна, b - емкость
    x = append(x, a) - для добавления элемента в срез
    element := x[index] - для доступа к конкретному элементу, index начинается с 0
    x[index] = element - для изменения элемента среза
    x = append(x[:index], x[index2]...) - для удаления элемента среза
    lenx = len(x) = для подсчета позиций в срезе