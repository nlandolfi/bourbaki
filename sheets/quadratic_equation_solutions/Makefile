# generated automatically by bbk sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make quadratic_equation_solutions.pdf

sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex

macros.tex: sheet.lit
	bbk macros sheet.lit > macros.tex

quadratic_equation_solutions.tex: sheet.tex macros.tex
	bbk sheet -mode c -in sheet.tex > quadratic_equation_solutions.tex

quadratic_equation_solutions.pdf: ../../*.tex ../../trademark.pdf *.tex quadratic_equation_solutions.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode quadratic_equation_solutions.tex
	../../latexrun quadratic_equation_solutions.tex
	make terms

graph.csv: sheet.tex
	bbk sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk graph --csv graph.csv --tmpl ../../bbk/static/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: quadratic_equation_solutions.pdf
	open quadratic_equation_solutions.pdf

o:
	make open

terms: quadratic_equation_solutions.tex
	bbk sheet -mode ts quadratic_equation_solutions.tex

spell:
	aspell -c sheet.tex

s:
	make spell

reset:
	rm -f graph.pdf
	rm -f graph.csv
	rm -f graph.graphviz

clean:
	make reset

remake:
	rm -f quadratic_equation_solutions.*
	make reset
	bbk sheet -mode mk > Makefile