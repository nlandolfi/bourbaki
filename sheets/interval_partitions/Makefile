# generated automatically by bbk sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make interval_partitions.pdf

sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex

macros.tex: sheet.lit
	bbk macros sheet.lit > macros.tex

interval_partitions.tex: sheet.tex macros.tex
	bbk sheet -mode c -in sheet.tex > interval_partitions.tex

interval_partitions.pdf: ../../*.tex ../../trademark.pdf *.tex interval_partitions.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode interval_partitions.tex
	../../latexrun interval_partitions.tex
	make terms

graph.csv: sheet.tex
	bbk sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk graph --csv graph.csv --tmpl ../../bbk/static/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: interval_partitions.pdf
	open interval_partitions.pdf

o:
	make open

terms: interval_partitions.tex
	bbk sheet -mode ts interval_partitions.tex

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
	rm -f interval_partitions.*
	make reset
	bbk sheet -mode mk > Makefile