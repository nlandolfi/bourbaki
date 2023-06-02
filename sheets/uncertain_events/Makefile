# generated automatically by bbk sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make uncertain_events.pdf

sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex

macros.tex: sheet.lit
	bbk macros sheet.lit > macros.tex

uncertain_events.tex: sheet.tex macros.tex
	bbk sheet -mode c -in sheet.tex > uncertain_events.tex

uncertain_events.pdf: ../../*.tex ../../trademark.pdf *.tex uncertain_events.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode uncertain_events.tex
	../../latexrun uncertain_events.tex
	make terms

graph.csv: sheet.tex
	bbk sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk graph --csv graph.csv --tmpl ../../bbk/static/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: uncertain_events.pdf
	open uncertain_events.pdf

o:
	make open

terms: uncertain_events.tex
	bbk sheet -mode ts uncertain_events.tex

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
	rm -f uncertain_events.*
	make reset
	bbk sheet -mode mk > Makefile