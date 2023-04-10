# generated automatically by bbk sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make equations.pdf

sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex

macros.tex: sheet.lit
	bbk macros sheet.lit > macros.tex

equations.tex: sheet.tex macros.tex
	bbk sheet -mode c -in sheet.tex > equations.tex

equations.pdf: ../../*.tex ../../trademark.pdf *.tex equations.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode equations.tex
	../../latexrun equations.tex
	make terms

graph.csv: sheet.tex
	bbk sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk graph --csv graph.csv --tmpl ../../bbk/static/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: equations.pdf
	open equations.pdf

o:
	make open

terms: equations.tex
	bbk sheet -mode ts equations.tex

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
	rm -f equations.*
	make reset
	bbk sheet -mode mk > Makefile