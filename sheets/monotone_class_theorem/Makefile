# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make monotone_class_theorem.pdf
	make graph.pdf

monotone_class_theorem.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > monotone_class_theorem.tex

monotone_class_theorem.pdf: ../../*.tex ../../trademark.pdf *.tex monotone_class_theorem.tex
	pdflatex --file-line-error -interaction=nonstopmode monotone_class_theorem.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: monotone_class_theorem.pdf
	open monotone_class_theorem.pdf

o:
	make open

terms: monotone_class_theorem.tex
	bbk_sheet -mode ts monotone_class_theorem.tex

spell:
	aspell -c monotone_class_theorem.tex

s:
	make spell

reset:
	rm -f graph.pdf
	rm -f graph.csv
	rm -f graph.graphviz

clean:
	make reset

remake:
	rm -f monotone_class_theorem.*
	make reset
	bbk_sheet -mode mk > Makefile