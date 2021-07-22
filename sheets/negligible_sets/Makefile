# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make negligible_sets.pdf

negligible_sets.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > negligible_sets.tex

negligible_sets.pdf: ../../*.tex ../../trademark.pdf *.tex negligible_sets.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode negligible_sets.tex
	../../latexrun negligible_sets.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: negligible_sets.pdf
	open negligible_sets.pdf

o:
	make open

terms: negligible_sets.tex
	bbk_sheet -mode ts negligible_sets.tex

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
	rm -f negligible_sets.*
	make reset
	bbk_sheet -mode mk > Makefile