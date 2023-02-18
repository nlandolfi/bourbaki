# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make relation_composites.pdf

sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex

relation_composites.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > relation_composites.tex

relation_composites.pdf: ../../*.tex ../../trademark.pdf *.tex relation_composites.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode relation_composites.tex
	../../latexrun relation_composites.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: relation_composites.pdf
	open relation_composites.pdf

o:
	make open

terms: relation_composites.tex
	bbk_sheet -mode ts relation_composites.tex

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
	rm -f relation_composites.*
	make reset
	bbk_sheet -mode mk > Makefile