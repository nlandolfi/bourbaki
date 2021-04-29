# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make finite_sets.pdf
	make graph.pdf

finite_sets.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > finite_sets.tex

finite_sets.pdf: ../../*.tex ../../trademark.pdf *.tex finite_sets.tex
	# pdflatex --file-line-error -interaction=nonstopmode finite_sets.tex
	../../latexrun finite_sets.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: finite_sets.pdf
	open finite_sets.pdf

o:
	make open

terms: finite_sets.tex
	bbk_sheet -mode ts finite_sets.tex

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
	rm -f finite_sets.*
	make reset
	bbk_sheet -mode mk > Makefile