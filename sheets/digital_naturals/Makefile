# generated automatically by bbk sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make digital_naturals.pdf

sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex

macros.tex: sheet.lit
	bbk macros sheet.lit > macros.tex

digital_naturals.tex: sheet.tex macros.tex
	bbk sheet -mode c -in sheet.tex > digital_naturals.tex

digital_naturals.pdf: ../../*.tex ../../trademark.pdf *.tex digital_naturals.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode digital_naturals.tex
	../../latexrun digital_naturals.tex
	make terms

graph.csv: sheet.tex
	bbk sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk graph --csv graph.csv --tmpl ../../bbk/static/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: digital_naturals.pdf
	open digital_naturals.pdf

o:
	make open

terms: digital_naturals.tex
	bbk sheet -mode ts digital_naturals.tex

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
	rm -f digital_naturals.*
	make reset
	bbk sheet -mode mk > Makefile