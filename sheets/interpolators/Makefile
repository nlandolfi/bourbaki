# generated automatically by bbk sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make interpolators.pdf

sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex

macros.tex: sheet.lit
	bbk macros sheet.lit > macros.tex

interpolators.tex: sheet.tex macros.tex
	bbk sheet -mode c -in sheet.tex > interpolators.tex

interpolators.pdf: ../../*.tex ../../trademark.pdf *.tex interpolators.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode interpolators.tex
	../../latexrun interpolators.tex
	make terms

graph.csv: sheet.tex
	bbk sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk graph --csv graph.csv --tmpl ../../bbk/static/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: interpolators.pdf
	open interpolators.pdf

o:
	make open

terms: interpolators.tex
	bbk sheet -mode ts interpolators.tex

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
	rm -f interpolators.*
	make reset
	bbk sheet -mode mk > Makefile