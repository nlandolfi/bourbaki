# generated automatically by bbk sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make threshold_graphs.pdf

sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex

macros.tex: sheet.lit
	bbk macros sheet.lit > macros.tex

threshold_graphs.tex: sheet.tex macros.tex
	bbk sheet -mode c -in sheet.tex > threshold_graphs.tex

threshold_graphs.pdf: ../../*.tex ../../trademark.pdf *.tex threshold_graphs.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode threshold_graphs.tex
	../../latexrun threshold_graphs.tex
	make terms

graph.csv: sheet.tex
	bbk sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk graph --csv graph.csv --tmpl ../../bbk/static/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: threshold_graphs.pdf
	open threshold_graphs.pdf

o:
	make open

terms: threshold_graphs.tex
	bbk sheet -mode ts threshold_graphs.tex

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
	rm -f threshold_graphs.*
	make reset
	bbk sheet -mode mk > Makefile