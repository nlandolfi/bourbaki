# generated automatically by bbk sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make signed_measure_examples.pdf

sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex

macros.tex: sheet.lit
	bbk macros sheet.lit > macros.tex

signed_measure_examples.tex: sheet.tex macros.tex
	bbk sheet -mode c -in sheet.tex > signed_measure_examples.tex

signed_measure_examples.pdf: ../../*.tex ../../trademark.pdf *.tex signed_measure_examples.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode signed_measure_examples.tex
	../../latexrun signed_measure_examples.tex
	make terms

graph.csv: sheet.tex
	bbk sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk graph --csv graph.csv --tmpl ../../bbk/static/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: signed_measure_examples.pdf
	open signed_measure_examples.pdf

o:
	make open

terms: signed_measure_examples.tex
	bbk sheet -mode ts signed_measure_examples.tex

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
	rm -f signed_measure_examples.*
	make reset
	bbk sheet -mode mk > Makefile