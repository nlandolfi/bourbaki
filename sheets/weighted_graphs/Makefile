# generated automatically by bbk sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make weighted_graphs.pdf

sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex

macros.tex: sheet.lit
	bbk macros sheet.lit > macros.tex

weighted_graphs.tex: sheet.tex macros.tex
	bbk sheet -mode c -in sheet.tex > weighted_graphs.tex

weighted_graphs.pdf: ../../*.tex ../../trademark.pdf *.tex weighted_graphs.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode weighted_graphs.tex
	../../latexrun weighted_graphs.tex
	make terms

graph.csv: sheet.tex
	bbk sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk graph --csv graph.csv --tmpl ../../bbk/static/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: weighted_graphs.pdf
	open weighted_graphs.pdf

o:
	make open

terms: weighted_graphs.tex
	bbk sheet -mode ts weighted_graphs.tex

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
	rm -f weighted_graphs.*
	make reset
	bbk sheet -mode mk > Makefile