# generated automatically by bbk sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make interval_graphs.pdf

sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex

macros.tex: sheet.lit
	bbk macros sheet.lit > macros.tex

interval_graphs.tex: sheet.tex macros.tex
	bbk sheet -mode c -in sheet.tex > interval_graphs.tex

interval_graphs.pdf: ../../*.tex ../../trademark.pdf *.tex interval_graphs.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode interval_graphs.tex
	../../latexrun interval_graphs.tex
	make terms

graph.csv: sheet.tex
	bbk sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk graph --csv graph.csv --tmpl ../../bbk/static/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: interval_graphs.pdf
	open interval_graphs.pdf

o:
	make open

terms: interval_graphs.tex
	bbk sheet -mode ts interval_graphs.tex

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
	rm -f interval_graphs.*
	make reset
	bbk sheet -mode mk > Makefile