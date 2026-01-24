# go-way

Изучение идиоматичного Go через конкурентные решения простых задач: горутины, каналы, паттерны fan-in/fan-out, worker pools и пайплайны.

Exploring idiomatic Go through concurrent solutions to simple problems: goroutines, channels, fan-in/fan-out patterns, worker pools, and pipelines.

---

## 🚀 Попробуй сам!

Синтаксис Go выглядит обманчиво просто. Но умелое использование горутин и каналов требует перестройки мышления программиста.

**Вызов:** возьми любую задачу из репозитория, прочитай условие и попробуй решить её самостоятельно через конкурентность. Затем сравни своё решение с предложенным. Какие паттерны ты использовал? Что можно улучшить?

Лучший способ освоить go-way — это писать код своими руками. Не просто читай — **создавай**.

## 🚀 Try It Yourself!

Go's syntax looks deceptively simple. But skillful use of goroutines and channels requires rewiring how programmers think.

**Challenge:** pick any problem from the repository, read the description, and try solving it concurrently on your own. Then compare your solution with the one provided. Which patterns did you use? What could be improved?

The best way to master go-way is to write code with your own hands. Don't just read—**create**.

---

## Зачем это нужно?
Go предлагает особенный путь мышления о конкурентности: *разделяй память через коммуникацию*.
Этот репозиторий — безопасная песочница для практики конкурентных паттернов на простых, хорошо знакомых задачах. Никакой предметной сложности — только чистая работа с горутинами и каналами.

## Why?
Go encourages a unique way of thinking about concurrency: *share memory by communicating*.
This repository is a safe sandbox for practicing concurrent patterns on simple, familiar problems. No domain complexity—just pure goroutines and channels.

---

## Что внутри?
Решения небольших задач программирования, реализованные через призму конкурентности Go:
- горутины и каналы (буферизованные и небуферизованные)
- fan-in / fan-out
- worker pools
- pipeline архитектура
- координация и синхронизация

**Важно:** многие задачи можно решить проще и последовательно. Здесь конкурентность используется **намеренно** — чтобы почувствовать, как ведут себя примитивы Go и как компонуются паттерны.

## What's Inside?
Solutions to small programming problems implemented through the lens of Go concurrency:
- goroutines and channels (buffered and unbuffered)
- fan-in / fan-out
- worker pools
- pipeline architecture
- coordination and synchronization

**Important:** many problems can be solved simpler and sequentially. Here, concurrency is used **intentionally**—to feel how Go primitives behave and how patterns compose.

---

## Фокус
Не на минимальный код, а на:
- ясное выражение конкурентности
- понимание компромиссов
- практику ключевых абстракций Go

## Focus
Not on minimal code, but on:
- clear expression of concurrency
- understanding trade-offs
- practicing core Go abstractions

---

## Что это такое в целом?
**Это:**
- учебный и исследовательский проект
- коллекция идиоматичных Go примеров
- демонстрация конкурентных паттернов

**Это НЕ:**
- утверждение, что конкурентность всегда лучше
- попытка чрезмерно оптимизировать тривиальные задачи
- проект, ориентированный на бенчмарки

## What This Is (and Isn't)
**This is:**
- a learning and exploration project
- a collection of idiomatic Go examples
- a showcase of concurrent patterns

**This is NOT:**
- a claim that concurrency is always better
- an attempt to over-optimize trivial problems
- a benchmark-focused project

---

## Откуда задачи?
Используются небольшие задачи (например, с Codewars) как удобные тестовые кейсы, а также задачи демонстрирующие преимущества go-way:
- чёткие условия
- знакомые ограничения
- фокус на дизайне решения, а не сложности проблемы

## Where Do Problems Come From?
Small problems (e.g., from Codewars) are used as convenient test cases and some real-world tasks:
- clear specifications
- familiar constraints
- focus on solution design, not problem complexity

---

## Дисклеймер
Конкурентность — это инструмол, а не цель.
Примеры приоритизируют ясность паттернов над краткостью и могут быть избыточными для реальных задач такого масштаба.

## Disclaimer
Concurrency is a tool, not a goal.
Examples prioritize pattern clarity over brevity and may be excessive for real-world problems of this scale.

---

## Tags

`golang` `go` `concurrency` `goroutines` `channels` `concurrent-programming` `go-patterns` `fan-in` `fan-out` `worker-pool` `pipeline` `go-concurrency` `idiomatic-go` `learning-go` `go-examples` `concurrent-patterns` `go-programming` `goroutine-patterns` `channel-patterns` `go-best-practices` `kata` `codewars`

## License

MIT