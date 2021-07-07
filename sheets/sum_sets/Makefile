# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make sum_sets.pdf

sum_sets.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > sum_sets.tex

sum_sets.pdf: ../../*.tex ../../trademark.pdf *.tex sum_sets.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode sum_sets.tex
	../../latexrun sum_sets.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: sum_sets.pdf
	open sum_sets.pdf

o:
	make open

terms: sum_sets.tex
	bbk_sheet -mode ts sum_sets.tex

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
	rm -f sum_sets.*
	make reset
	bbk_sheet -mode mk > Makefile