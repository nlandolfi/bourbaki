# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make weighted_norms.pdf

weighted_norms.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > weighted_norms.tex

weighted_norms.pdf: ../../*.tex ../../trademark.pdf *.tex weighted_norms.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode weighted_norms.tex
	../../latexrun weighted_norms.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: weighted_norms.pdf
	open weighted_norms.pdf

o:
	make open

terms: weighted_norms.tex
	bbk_sheet -mode ts weighted_norms.tex

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
	rm -f weighted_norms.*
	make reset
	bbk_sheet -mode mk > Makefile