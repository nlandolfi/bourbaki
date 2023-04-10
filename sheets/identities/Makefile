# generated automatically by bbk sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make identities.pdf

sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex

macros.tex: sheet.lit
	bbk macros sheet.lit > macros.tex

identities.tex: sheet.tex macros.tex
	bbk sheet -mode c -in sheet.tex > identities.tex

identities.pdf: ../../*.tex ../../trademark.pdf *.tex identities.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode identities.tex
	../../latexrun identities.tex
	make terms

graph.csv: sheet.tex
	bbk sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk graph --csv graph.csv --tmpl ../../bbk/static/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: identities.pdf
	open identities.pdf

o:
	make open

terms: identities.tex
	bbk sheet -mode ts identities.tex

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
	rm -f identities.*
	make reset
	bbk sheet -mode mk > Makefile