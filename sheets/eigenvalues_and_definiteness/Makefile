# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make eigenvalues_and_definiteness.pdf

eigenvalues_and_definiteness.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > eigenvalues_and_definiteness.tex

eigenvalues_and_definiteness.pdf: ../../*.tex ../../trademark.pdf *.tex eigenvalues_and_definiteness.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode eigenvalues_and_definiteness.tex
	../../latexrun eigenvalues_and_definiteness.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: eigenvalues_and_definiteness.pdf
	open eigenvalues_and_definiteness.pdf

o:
	make open

terms: eigenvalues_and_definiteness.tex
	bbk_sheet -mode ts eigenvalues_and_definiteness.tex

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
	rm -f eigenvalues_and_definiteness.*
	make reset
	bbk_sheet -mode mk > Makefile