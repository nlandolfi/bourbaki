# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make ring_ideals.pdf

sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex

macros.tex: sheet.lit
	bbk macros sheet.lit > macros.tex

ring_ideals.tex: sheet.tex macros.tex
	bbk_sheet -mode c -in sheet.tex > ring_ideals.tex

ring_ideals.pdf: ../../*.tex ../../trademark.pdf *.tex ring_ideals.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode ring_ideals.tex
	../../latexrun ring_ideals.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: ring_ideals.pdf
	open ring_ideals.pdf

o:
	make open

terms: ring_ideals.tex
	bbk_sheet -mode ts ring_ideals.tex

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
	rm -f ring_ideals.*
	make reset
	bbk_sheet -mode mk > Makefile