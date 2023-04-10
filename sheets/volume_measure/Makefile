# generated automatically by bbk sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make volume_measure.pdf

sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex

macros.tex: sheet.lit
	bbk macros sheet.lit > macros.tex

volume_measure.tex: sheet.tex macros.tex
	bbk sheet -mode c -in sheet.tex > volume_measure.tex

volume_measure.pdf: ../../*.tex ../../trademark.pdf *.tex volume_measure.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode volume_measure.tex
	../../latexrun volume_measure.tex
	make terms

graph.csv: sheet.tex
	bbk sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk graph --csv graph.csv --tmpl ../../bbk/static/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: volume_measure.pdf
	open volume_measure.pdf

o:
	make open

terms: volume_measure.tex
	bbk sheet -mode ts volume_measure.tex

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
	rm -f volume_measure.*
	make reset
	bbk sheet -mode mk > Makefile