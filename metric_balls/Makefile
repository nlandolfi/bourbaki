metric_balls.pdf: ../*.tex ../trademark.pdf *.tex
	pdflatex metric_balls.tex
	make defs

make:
	make metric_balls.pdf

open: metric_balls.pdf
	open metric_balls.pdf

o: open
	make open

defs: metric_balls.tex
	defs metric_balls.tex

clean:
	rm *.aux
	rm *.out
	rm *.log
