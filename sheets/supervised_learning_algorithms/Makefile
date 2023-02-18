# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make supervised_learning_algorithms.pdf

sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex

supervised_learning_algorithms.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > supervised_learning_algorithms.tex

supervised_learning_algorithms.pdf: ../../*.tex ../../trademark.pdf *.tex supervised_learning_algorithms.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode supervised_learning_algorithms.tex
	../../latexrun supervised_learning_algorithms.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: supervised_learning_algorithms.pdf
	open supervised_learning_algorithms.pdf

o:
	make open

terms: supervised_learning_algorithms.tex
	bbk_sheet -mode ts supervised_learning_algorithms.tex

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
	rm -f supervised_learning_algorithms.*
	make reset
	bbk_sheet -mode mk > Makefile