# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make real_open_sets.pdf

real_open_sets.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > real_open_sets.tex

real_open_sets.pdf: ../../*.tex ../../trademark.pdf *.tex real_open_sets.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode real_open_sets.tex
	../../latexrun real_open_sets.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: real_open_sets.pdf
	open real_open_sets.pdf

o:
	make open

terms: real_open_sets.tex
	bbk_sheet -mode ts real_open_sets.tex

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
	rm -f real_open_sets.*
	make reset
	bbk_sheet -mode mk > Makefile